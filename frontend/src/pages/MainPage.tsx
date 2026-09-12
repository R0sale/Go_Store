import { useLoaderData } from "react-router"
import type { item } from "../models/item"
import { useState } from "react"
import { Catalog } from "../components/Catalog"
import { Header } from "../components/Header"
import { newApiService } from "../services/api_service"
import ErrorBoundary from "../components/ErrorBoundary/ErrorBoundary"


export const loadCatalogItems = async (): Promise<item[] | null> => {
    const apiService = newApiService()

    const { ok, response } = await apiService.apiCall(import.meta.env.VITE_CATALOG_URL + "/api/catalog", "GET")

    if (!ok) {
        throw new Error("Couldn't load Catalog items for main page")
    }

    return response as item[]
}

export const MainPage = () => {
    const items = useLoaderData<item[] | null>()
    const [catalogItems, setCatalogItems] = useState<item[] | null>(items)

    return (
        <ErrorBoundary fallback={<p>Sorry</p>}>
            <div className="min-h-screen bg-slate-50">
                <Header />
                <ErrorBoundary fallback={<p>Sorry, cant load anything</p>}>
                    <Catalog items={catalogItems}/>
                </ErrorBoundary>
            </div>
        </ErrorBoundary>
    )
}