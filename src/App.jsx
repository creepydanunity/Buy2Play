import './App.scss'
import Home from './pages/home/Home'
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Products from './pages/products/Products';
import Chat from './pages/chat/Chat';
import Anouncement from './components/anouncement/Anouncement';
import Authorization from "./pages/authorization/Authorization.jsx";

function App() {
  return (
    <BrowserRouter>
      <Anouncement/>
      <Routes>
          <Route  path="/" element={<Home/>} />
          <Route  path="/products" element={<Products/>} />
          <Route  path="/authorization" element={<Authorization/>} />
          <Route  path="/chat" element={<Chat/>} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
