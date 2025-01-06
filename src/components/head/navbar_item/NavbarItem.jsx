import React, { useState, useRef, useEffect} from 'react';
import "./NavbarItem.scss";


function NavbarItem({ item, index }) {
  const [isOpen, setIsOpen] = useState(false);
  const [textHeight, setTextHeight] = useState('0px');
  const textRef = useRef(null);
  const containerRef = useRef(null);

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
    <div className='navbar-item-container' ref={containerRef}>
      <div className='header-navbar-item' onClick={toggleText}>
        {item.title}
        <img className={`item-arrow ${isOpen ? 'rotate' : ''}`} src={item.arrow} alt="arrow" />
      </div>
      <ul
        className={`navbar-text ${isOpen ? 'open' : ''}`}
        style={{ height: textHeight }}
        ref={textRef}
      >
          {item.names.map((name, index) => (
             <li key={index}>{name}</li>
          ))}
      </ul>
    </div>
  );
}

export default NavbarItem;
