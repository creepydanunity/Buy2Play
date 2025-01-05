import "./OurPreferences.scss"

function OurPreferences(){
    return(
    <div className='our-preferences-section'>
            <div className='preferences-section-wrapper'>
                <h2 className="preferences-section-title">Наши преимущества</h2>
                <ul className="preferences-list">
                    <li className="preferences-list-item">
                        <div>
                            <img src="lightning.svg" alt="lightning" />
                            Скорость доставки
                        </div>
                        <img src="arrow.svg" alt="arrow" />
                    </li>
                    <li className="preferences-list-item">
                        <div>
                            <img src="lock.svg" alt="lock" />
                            Безопасность
                        </div>
                        <img src="arrow.svg" alt="arrow" />
                    </li>
                    <li className="preferences-list-item">
                        
                        <div>
                            <img src="box.svg" alt="box" />
                            Доступность
                        </div>
                        <img src="arrow.svg" alt="arrow" />
                    </li>
                    <li className="preferences-list-item">
                        
                        <div>
                            <img src="star.svg" alt="star" />
                            Конкурсы
                        </div>
                        <img src="arrow.svg" alt="arrow" />
                    </li>
                </ul>
            </div>

    </div>

    )
}

export default OurPreferences;