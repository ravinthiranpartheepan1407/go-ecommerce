import Image from "next/image"
import {StarIcon} from "@heroicons/react/solid";

export default function Product({id, title, price, description, category, image, rating}){
    return(
        <div>
            <br />
            <br />
            <div class="max-w-sm rounded overflow-hidden shadow-lg">
            <img class="w-full" src={image} alt="Go-Ecommerce" />
            <div class="px-6 py-4">
                <div class="font-bold text-xl mb-2">{title}</div>
                <p class="text-gray-700 text-base">
                    {description}
                </p>
            </div>
            <div class="px-6 pt-4 pb-2">
                <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">{category}</span>
     
                <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">{price}</span>
            </div>
            <button className="bg-black w-full p-4 text-white">Buy</button>
            </div>
            <br />
            <br />
        </div>
    )
}