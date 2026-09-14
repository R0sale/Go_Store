import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { loadCatalogItems } from './pages/MainPage.tsx'
import { MainPage } from './pages/MainPage.tsx'
import { createBrowserRouter, RouterProvider } from "react-router"
import { Fallback } from './components/ErrorBoundary/GlobalFallback.tsx'
import { LoginPage } from './pages/LoginPage.tsx'
import { RegisterPage } from './pages/RegisterPage.tsx'

const router = createBrowserRouter([
  {
    path: "/",
    loader: loadCatalogItems,
    Component: MainPage,
    errorElement: <Fallback />
  },
  {
    path: "/login",
    Component: LoginPage
  },
  {
    path: "/register",
    Component: RegisterPage
  }
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
