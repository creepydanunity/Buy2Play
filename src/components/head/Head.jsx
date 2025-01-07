import "./Head.scss"
import NavbarItem from "./navbar_item/NavbarItem";
import preferences from "./navbar_item/item_data";
import SearchBar from "../search_bar/SearchBar.jsx";
import {useState} from "react";
import { Link } from "react-router-dom";

function Head(){
    const [isSearchOpen, setIsSearchOpen] = useState(false);

    const toggleSearchBar = () => {
        setIsSearchOpen(!isSearchOpen);
    };

    return(
    <header className="header">
        <div className="header-wrapper">
            <div className='header-user-actions'>
                <Link to="/">
                    <img src="logo.svg" width={100} height={48} alt="B2P" className="a"/>
                </Link>
                <div className='user-actions-controls'>
                    <img src="search.svg" width={20} height={20} alt="magnifier" onClick={() => setIsSearchOpen(!isSearchOpen)}
                    />
                    <Link to="/authorization">
                        <img src="profile.svg" width={20} height={20} alt="profile"/>
                    </Link>
                    <img src="shopping-bag.svg" width={20} height={20} alt="shopping-bag"/>
                </div>
            </div>

            {isSearchOpen && <SearchBar onClose={toggleSearchBar} />}

            <div className='header-navbar'>
                <Link to="/" className='header-navbar-item'>Главная</Link>
                {preferences.map(item => (
                    <NavbarItem key={item.id} item={item}/>
                ))}
                <Link to="/" className='header-navbar-item'>Spotify</Link>
            </div>

        </div>

    </header>
    )
}

export default Head;