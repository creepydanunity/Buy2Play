import "./Head.scss"
import NavbarItem from "./navbar_item/NavbarItem";
import preferences from "./navbar_item/item_data";
import SearchBar from "../search_bar/SearchBar.jsx";
import {useState} from "react";


function Head(){
    const [isSearchOpen, setIsSearchOpen] = useState(false); // Состояние для управления поисковой строкой

    const toggleSearchBar = () => {
        setIsSearchOpen(!isSearchOpen);
    };

    return(
    <header className="header">
        <div className="header-wrapper">
            <div className='header-user-actions'>
                <img src="logo.svg" width={100} height={48} alt="B2P" className="a"/>
                <div className='user-actions-controls'>
                    <img src="search.svg" width={20} height={20} alt="magnifier" onClick={() => setIsSearchOpen(!isSearchOpen)}
                    />
                    <a href="/authorization">
                        <img src="profile.svg" width={20} height={20} alt="profile"/>
                    </a>
                    <img src="shopping-bag.svg" width={20} height={20} alt="shopping-bag"/>
                </div>
            </div>

            {isSearchOpen && <SearchBar onClose={toggleSearchBar} />}

            <ul className='header-navbar'>
                <li className='header-navbar-item'>Главная</li>
                {preferences.map(item => (
                    <NavbarItem key={item.id} item={item}/>
                ))}
                <li className='header-navbar-item'>Spotify</li>
            </ul>

        </div>

    </header>
    )
}

export default Head;