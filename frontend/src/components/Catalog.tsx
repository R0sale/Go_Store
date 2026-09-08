import { useState } from "react"
import type { item } from "../models/item"
import { useLoaderData } from "react-router"

export const loadCatalogItems = async (): Promise<item[] | null> => {
    const response = await fetch("", {
        method: "GET",
        credentials: "include"
    })

    if (!response.ok) {
        return null
    }

    const result = await response.json() as item[]

    return result
}

export const Catalog = () => {
    const { items } = useLoaderData()
    const [items, setItems] = useState<item[]>()

    return (
        <div>

        </div>
    )
}