import "./Games.scss"
import GamesCard from "./GamesCard/GamesCard";
import cards from "./GamesCard/cards_data";

function Games(){
    return(
    <section className='popular-games-section'>
        <div className='popular-games-wrapper'>
            <h2>Популярные игры</h2>
            <div className='games-card-wrapper'>
                {cards.map((card) => (
                        <GamesCard key={card.id} {...card} />
                    ))}
            </div>
        </div>
    </section>
    )
}

export default Games;