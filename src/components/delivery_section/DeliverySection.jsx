import "./DeliverySection.scss"
import IMAGES from "../../images/Images";


function DeliverySection(){
    return(
        <div className='delivery-section'>
            <div className='delivery-section-wrapper'>
                <img className="delivery-section-image" src={IMAGES.delivery}alt="rocket" />
                <div className='delivery-section-content'>
                    <h1 className='delivery-content-title'>Моментальная доставка</h1>
                    <div className='delivery-content-description'>
                    Большинство товаров в нашем магазине имеют
                    моментальную доставку. Всё просто! Вы
                    оплачиваете понравившийся товар и получаете
                    его на указанную электронную почту. Не нужно
                    ждать ответа продавца или следить за графиком
                    работы магазина, в считанные минуты.
                    </div>
                    <button className='delivery-content-button'>К покупкам</button>
                </div>
            </div>

        </div>
    )
}

export default DeliverySection;