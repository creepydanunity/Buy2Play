import "./Products.scss"
import Head from "../../components/head/Head";
import Foot from "../../components/foot/Foot";
import ProductsHeroSection from "../../components/products_hero_section/ProductsHeroSection";
import ProductList from "../../components/product_list/ProductList";
import ProductSort from "../../components/products_sort/ProductSort";


function Products(){
    return(
        <div>
            <Head/>
            <ProductsHeroSection/>
            <ProductList/>
            <Foot/>
        </div>
    )
}

export default Products;