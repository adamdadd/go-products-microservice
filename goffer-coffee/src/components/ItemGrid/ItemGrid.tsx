import React from "react";
import { ItemCard, ItemGridStyle } from "./ItemGrid.styles";

export type Product = {
    id: number;
    name: string;
    description: string;
    price: number;
    sku: string;
    image_url: string;
}


type props = {
    productList: Array<Product>;
}

export const ItemGrid = ({ productList } : props) => {
  return (
    <ItemGridStyle>
      {productList.map((product : Product) => {
        <ItemCard> 
            <h3> {product.name} </h3>
        </ItemCard>;
      })}
    </ItemGridStyle>
  );
}
