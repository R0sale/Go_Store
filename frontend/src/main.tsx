import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { loadCatalogItems } from './pages/MainPage.tsx'
import { MainPage } from './pages/MainPage.tsx'
import { createBrowserRouter, RouterProvider } from "react-router"
import { Fallback } from './components/ErrorBoundary/GlobalFallback.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    loader: loadCatalogItems,
    Component: MainPage,
    errorElement: <Fallback />
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
