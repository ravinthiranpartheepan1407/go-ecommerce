import Product from "./Product";

export default function ProductContainer({productLists}){
    return(
        <div className="lg:grid grid-cols-4 gap-4 mx-auto">
            {productLists.slice(0,productLists.length).map(({id, title, price, description, category, image})=>(
            <Product
                key={id}
                id={id}
                title={title}
                price={price}
                category={category}
                image={image}
                rating={5}
            />
            ))}
        </div>
    )
}