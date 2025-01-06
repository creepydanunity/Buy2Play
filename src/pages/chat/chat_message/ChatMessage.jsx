import "./ChatMessage.scss";
import React from 'react';

function ChatMessage({ message, timestamp, user, isCurrentUser }) {
  const messageClass = isCurrentUser ? "message user" : "message other";

  return (
    <div className={messageClass}>
        <span className="timestamp">{timestamp}</span>
        {user && <span className="user-name">{user}</span>} 
        <p>{message}</p>

    </div>
  );
}

export default ChatMessage;

