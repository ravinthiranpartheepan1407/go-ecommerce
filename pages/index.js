import styles from "../styles/Home.module.css";
import Header from "../components/header";
import ProductContainer from "../components/ProductContainer";
import {goEcommerceDataset} from "/public/dataset";


export default function Home({products}) {
    return(
      <div>
        <Header />
        <div>
          <main className="max-w-screen-2xl mx-auto">
           <ProductContainer productLists={products} />
          </main>
        </div>
      </div>
    );
}

export async function getServerSideProps(context){
  const products = goEcommerceDataset
    return{
      props:{
        products,
      },
    };
  }