import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import { Header } from './components/Header/Header';
import { ItemGrid, Product } from './components/ItemGrid/ItemGrid';
import './App.css';

function App() {
  const [productList, setProductList] = useState<Array<Product>>([])

  useEffect(() => {
    fetch("/products/", {method: "GET"})
    .then( response => { if (!response.ok) {
        throw new Error(response.statusText)
      }
      return response.json()
    })
    .then(data => setProductList(data))
  }, [])

  return (
    <div>
      <Header />
      <ItemGrid productList={productList}/>
    </div>
  );
}

export default App;
