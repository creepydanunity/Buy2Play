import "./PreferencesItem.scss";
import React, { useState, useRef, useEffect } from 'react';

function PreferencesItem({ id, title, text, image, arrow }) {
    const [isOpen, setIsOpen] = useState(false);
    const [textHeight, setTextHeight] = useState('0px');
    const textRef = useRef(null);

    const toggleText = () => {
        setIsOpen(!isOpen);
    };

    useEffect(() => {
        if (isOpen && textRef.current) {
            setTextHeight(`${textRef.current.scrollHeight}px`);
        } else {
           setTextHeight('0px');
        }
    }, [isOpen, textRef.current]);

    return (
        <div>
            <div className="preferences-list-item" onClick={toggleText}>
                <div className="preferneces-img-wrapper">
                    <img className="preferneces-img" src={image} alt="lightning" />
                    {title}
                </div>
                
                <img className={`preferences-arrow ${isOpen ? 'rotate' : ''}`} src={arrow} alt="arrow"/>
            </div>

            <div 
              className={`preferences-text ${isOpen ? 'open' : ''}`}
              style={{ height: textHeight }}
              ref={textRef}
            >
                {text}
            </div>
        </div>
    );
}

export default PreferencesItem;
    