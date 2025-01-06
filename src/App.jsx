import './App.scss'
import Home from './pages/home/Home'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/products/Products';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route  path="/" element={<Home/>} />
        <Route  path="/products" element={<Products/>} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
