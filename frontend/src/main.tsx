import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { App } from './App.tsx'
import { loadCatalogItems } from './components/Catalog.tsx'
import { MainPage } from './pages/MainPage.tsx'
import { createBrowserRouter } from "react-router"

createBrowserRouter([
  {
    path: "/",
    loader: loadCatalogItems,
    Component: MainPage
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
