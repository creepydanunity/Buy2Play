import "./Head.scss"
import NavbarItem from "./navbar_item/NavbarItem";
import preferences from "./navbar_item/item_data";


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
            <ul className='header-navbar'>
                <li className='header-navbar-item'>Главная</li>
                    {preferences.map(item => ( 
                <NavbarItem key={item.id} item={item} /> 
                ))}
                <li className='header-navbar-item'>Spotify</li>
            </ul>
    
        </div>
        
    </header>
    )
}

export default Head;