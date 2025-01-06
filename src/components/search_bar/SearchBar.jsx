import { useState } from "react";
import "./SearchBar.scss";

// eslint-disable-next-line react/prop-types
function SearchBar({ onClose }) {
    const [query, setQuery] = useState("");

    const handleSearch = (e) => {
        e.preventDefault();
        console.log("Search query:", query);
    };

    return (
        <div className="search-bar-container">
            <form className="search-form" onSubmit={handleSearch}>
                <input
                    type="text"
                    placeholder="Введите запрос"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    className="search-input"
                    autoFocus
                />
                <button type="submit" className="search-button">
                    <img src="search.svg" width={20} height={20} alt="magnifier"/>
                </button>
            </form>
            <button type="button" className="close-button" onClick={onClose}>
                &#x2715;
            </button>
        </div>
    );
}

export default SearchBar;
