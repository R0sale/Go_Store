import { useLoaderData } from "react-router"
import type { item } from "../models/item"
import { useState } from "react"
import { Catalog } from "../components/Catalog"
import { newApiService } from "../services/api_service"
import { Layout } from "../components/Layout"


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
        <Layout>
            <Catalog items={catalogItems}/>
        </Layout>
    )
}