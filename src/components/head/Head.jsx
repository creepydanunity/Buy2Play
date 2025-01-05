import "./Head.scss"

function Head(){
    return(
    <header className="header">
        <div className="header-wrapper">
            <div className='header-user-actions'>
                <img src="logo.svg" width={100} height={48} alt="B2P" className="a"/>
                <div className='user-actions-controls'>
                    <img src="search.svg" width={20} height={20} alt="magnifier"/>
                    <img src="profile.svg" width={20} height={20} alt="profile" />
                    <img src="shopping-bag.svg" width={20} height={20}alt="shopping-bag" />
                </div>
            </div>
            <div className='header-navbar'>
                <ul className='header-navbar-row'>
                    <li>Главная</li>
                    <li>Мобильные игры</li>
                    <li>Мобильные игры</li>
                    <li>Игровая валюта</li>
                    <li>Spotify</li>
                </ul>
            </div> 
        </div>
        
    </header>
    )
}

export default Head;