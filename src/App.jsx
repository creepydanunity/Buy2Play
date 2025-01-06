import './App.scss'
import Home from './pages/home/Home'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/products/Products';
import Chat from './pages/chat/Chat';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route  path="/" element={<Home/>} />
        <Route  path="/products" element={<Products/>} />
        <Route  path="/chat" element={<Chat/>} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
