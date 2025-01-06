import "./GamesCard.scss"

function GamesCard({ id, image, title, subtitle, buttonLabel }){
    return(
        <div className='popular-games-card'>
            <div className="card-img-wrapper">
                <img src={image} alt={title} />
            </div>
            
            <h3 className="games-card-title">{title}</h3>
            <div className="games-card-text">{subtitle}</div>
            <a href="" className="games-card-link">{buttonLabel}</a>
        </div>
    )
}

export default GamesCard;