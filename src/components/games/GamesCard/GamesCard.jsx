import "./GamesCard.scss"
import IMAGES from "../../../images/Images";

function GamesCard(){
    return(
        <div className='popular-games-card'>
            <img src={IMAGES.roblox_game} alt="Roblox" />
            <h3 className="games-card-title">Roblox</h3>
            <div className="games-card-text">
            Подарочные карты, пополнение с
            заходом на аккаунт или может быть
            геймпасс? У нас вы найдете любой
            метод пополнения робуксов по самым низким ценам!
            </div>
            <button className="games-card-button">Купить робуксы</button>
        </div>
    )
}

export default GamesCard;