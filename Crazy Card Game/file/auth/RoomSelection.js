import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./RoomSelection.module.css"; // Importing CSS as a module

function RoomSelection() {
  const navigate = useNavigate();
  const [rooms, setRooms] = useState([]);
  const [openingRoom, setOpeningRoom] = useState(null);

  useEffect(() => {
    // Simulating fetching room data, with random availability (open/closed)
    const fetchedRooms = Array.from({ length: 4 }, (_, i) => ({
      number: i + 1,
      suit: ["hearts", "diamonds", "clubs", "spades"][i],
      symbol: ["♥", "♦", "♣", "♠"][i],
      isOpen: Math.random() > 0.5,  // Random open/closed rooms
    }));
    setRooms(fetchedRooms);
  }, []);

  const enterRoom = (room) => {
    if (!room.isOpen) {
      alert("Room is full! Choose another.");
      return;
    }

    setOpeningRoom(room.number);  // Mark the room as opening
    setTimeout(() => {
      navigate(`/crazycardgame/${room.number}`); // Navigate to the game page
    }, 500);
  };

  return (
    <div className={styles.roomContainer}>
      <h1>Select a Room</h1>
      <div className={styles.roomList}>
        {rooms.map((room) => (
          <div
            key={room.number}
            className={`${styles.room} 
                        ${room.isOpen ? styles.openRoom : styles.closedRoom} 
                        ${openingRoom === room.number ? styles.opening : ""} 
                        ${styles[room.suit]}`} // Applying the suit as a class
            onClick={() => enterRoom(room)} // Navigate to the game when clicked
          >
            <span className={styles.symbol}>{room.symbol}</span>
            <span className={styles.roomText}>{room.suit.charAt(0).toUpperCase() + room.suit.slice(1)} Room</span>
          </div>
        ))}
      </div>
    </div>
  );
}

export default RoomSelection;
