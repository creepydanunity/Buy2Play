import "./Games.scss"
import GamesCard from "./GamesCard/GamesCard";


function Games(){
    return(
    <div className='popular-games-section'>
        <div className='popular-games-wrapper'>
            <h2>Популярные игры</h2>
            <div className='games-card-wrapper'>
                <GamesCard/>
                <GamesCard/>
                <GamesCard/>
            </div>
        </div>
    </div>
    )
}

export default Games;