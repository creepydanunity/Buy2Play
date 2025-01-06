import "./Chat.scss";
import ChatMessage from "./chat_message/ChatMessage";
import React, { useState, useRef, useEffect } from 'react';

function Chat() {
    const [messages, setMessages] = useState([]);
    const [newMessageText, setNewMessageText] = useState('');
    const chatMessagesRef = useRef(null);

    const handleSendMessage = () => {
        if (newMessageText.trim() !== '') {
            const newMessage = {
                message: newMessageText,
                timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
                user: 'You',
                isCurrentUser: true,
            };

            setMessages(prevMessages => [...prevMessages, newMessage]);
            setNewMessageText('');
            

            setTimeout(() => {
                const botMessage = {
                    message: "Отлично!",
                    timestamp: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
                    user: 'Bot',
                    isCurrentUser: false,
                }
                setMessages(prevMessages => [...prevMessages, botMessage]);

            }, 1000)
        }
    };

    const handleInputChange = (e) => {
        setNewMessageText(e.target.value);
    };

    const handleKeyDown = (e) => {
        if(e.key === 'Enter'){
            handleSendMessage();
        }
    }


    useEffect(() => {
        if (chatMessagesRef.current) {
            chatMessagesRef.current.scrollTop = chatMessagesRef.current.scrollHeight;
        }
    }, [messages]);

    return (
        <div className="chat-container">
            <header className="chat-header">
                <div className="chat-icon">C</div>
                <div className="chat-title">
                    <h1>ChatFlow</h1>
                    <p>
                        A live chat interface that allows for seamless, natural communication
                        and connection.
                    </p>
                </div>
                <button className="close-button">&times;</button>
            </header>

            <div className="chat-messages" ref={chatMessagesRef}>
                {messages.map((msg, index) => (
                    <ChatMessage key={index} {...msg} />
                ))}
            </div>


            <div className="chat-input">
                <input 
                    type="text" 
                    placeholder="Reply ..." 
                    value={newMessageText}
                    onChange={handleInputChange}
                    onKeyDown={handleKeyDown}
                />
                <button className="send-button" onClick={handleSendMessage}>&#9658;</button>
            </div>
        </div>
    );
}

export default Chat;
