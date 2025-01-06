import "./Home.scss"
import Head from "../../components/head/Head";
import HeroSection from "../../components/hero_section/HeroSection";
import Foot from "../../components/foot/Foot";
import DeliverySection from "../../components/delivery_section/DeliverySection";
import Games from "../../components/games/Games";
import OurPreferences from "../../components/our_preferences/OurPreferences";


function Home(){
    return(
    <div className='home-wrapper'>
        <Head/>
        <HeroSection/>
     

        <DeliverySection/>

        <Games/>

        <OurPreferences/>

        <Foot/>
 
    </div>
    )
}

export default Home;