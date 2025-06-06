'use strict';

const startButton = document.getElementById('startButton');
const hangupButton = document.getElementById('hangupButton');
hangupButton.disabled = true;

const localVideo = document.getElementById('localVideo');
const remoteVideo = document.getElementById('remoteVideo');

let pc;
let localStream;
let socket;

// Connect to the WebSocket server
socket = new WebSocket('ws://192.168.100.5:8080/ws');
console.log(socket)

socket.onopen = () => {
  console.log('Connected to WebSocket server');
};

socket.onmessage = (event) => {
  const message = JSON.parse(event.data);

  if (!localStream) {
    console.log('Not ready yet');
    return;
  }

  switch (message.type) {
    case 'offer':
      handleOffer(message);
      break;
    case 'answer':
      handleAnswer(message);
      break;
    case 'ice-candidate':
      handleCandidate(message);
      break;
    // case 'ready':
    //   if (pc) {
    //     console.log('Already in call, ignoring');
    //     return;
    //   }
    //   makeCall();
    //   break;
    // case 'bye':
    //   if (pc) {
    //     hangup();
    //   }
    //   break;
    default:
      console.log('Unhandled message:', message);
      break;
  }
};

socket.onerror = (error) => {
  console.error('WebSocket error:', error);
};

socket.onclose = () => {
  console.log('WebSocket connection closed');
};

startButton.onclick = async () => {
  localStream = await navigator.mediaDevices.getUserMedia({ audio: true, video: true });
  localVideo.srcObject = localStream;

  startButton.disabled = true;
  hangupButton.disabled = false;

  sendSignal({ type: 'ready' });
};

hangupButton.onclick = async () => {
  hangup();
  sendSignal({ type: 'bye' });
};

async function hangup() {
  if (pc) {
    pc.close();
    pc = null;
  }
  localStream.getTracks().forEach((track) => track.stop());
  localStream = null;
  startButton.disabled = false;
  hangupButton.disabled = true;
}

function createPeerConnection() {
  pc = new RTCPeerConnection();
  pc.onicecandidate = (e) => {
    const message = {
      type: 'candidate',
      candidate: null,
    };
    if (e.candidate) {
      message.candidate = e.candidate.candidate;
      message.sdpMid = e.candidate.sdpMid;
      message.sdpMLineIndex = e.candidate.sdpMLineIndex;
    }
    sendSignal(message);
  };
  pc.ontrack = (e) => (remoteVideo.srcObject = e.streams[0]);
  localStream.getTracks().forEach((track) => pc.addTrack(track, localStream));
}

async function makeCall() {
  await createPeerConnection();

  const offer = await pc.createOffer();
  await pc.setLocalDescription(offer);

  sendSignal({ type: 'offer', sdp: offer.sdp });
}

async function handleOffer(offer) {
  if (pc) {
    console.error('Existing peer connection');
    return;
  }
  await createPeerConnection();
  await pc.setRemoteDescription(new RTCSessionDescription(offer));

  const answer = await pc.createAnswer();
  await pc.setLocalDescription(answer);

  sendSignal({ type: 'answer', sdp: answer.sdp });
}

async function handleAnswer(answer) {
  if (!pc) {
    console.error('No peer connection');
    return;
  }
  await pc.setRemoteDescription(new RTCSessionDescription(answer));
}

async function handleCandidate(candidate) {
  if (!pc) {
    console.error('No peer connection');
    return;
  }
  if (!candidate.candidate) {
    await pc.addIceCandidate(null);
  } else {
    await pc.addIceCandidate(new RTCIceCandidate(candidate));
  }
}

function sendSignal(message) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message));
  } else {
    console.error('WebSocket is not open');
  }
}
