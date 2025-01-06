import React, { useEffect, useState } from "react";
import data from "./data";
import "./Slider.scss";


function Slider() {
    const [people, setPeople] = useState(data);
    const [currentIndex, setIndex] = useState(0);

    useEffect(() => {
        const lastIndex = people.length - 1;
        if (currentIndex < 0) {
            setIndex(lastIndex);
        }
        if (currentIndex > lastIndex) {
            setIndex(0);
        }
    }, [currentIndex, people]);

    return (
        <section className="slider">
            <div className="slider-wrapper">
                {people.map((person, personIndex) => {
                    const { id, image, subtitle, title, buttonLabel } = person;
                    let position = "nextSlide";
                    if (personIndex === currentIndex) {
                        position = "activeSlide";
                    }

                    if (personIndex === currentIndex - 1 || (currentIndex === 0 && personIndex === people.length - 1)) {
                        position = "lastSlide";
                    }
                    return (
                        
                            <article 
                                className={position} 
                                key={id}
                                style={{ backgroundImage: `url(${image})` }}
                            >
                            <div className="slider-text">
                                <div className="slider-text-wrapper">
                                    <h2 className="slider-header">{title}</h2>
                                    <p className="slider-subtitle">{subtitle}</p>
                                    <div className="slider-button-wrapper">
                                        <button className="slider-link-button">{buttonLabel}</button>
                                    </div>
                                    
                                </div>
               
                            </div>
                            </article>
                        

                    );
                })}

            </div>
            <div className="controls">
                <button className="prev" onClick={() => setIndex(prev => prev - 1)}>
                        {"<"}
                    </button>
                <button className="next" onClick={() => setIndex(prev => prev + 1)}>
                        {">"}
                </button>
            </div>

        </section>
    );
}

export default Slider;
