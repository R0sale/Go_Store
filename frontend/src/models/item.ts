interface catalogItem {
    id: number;
    name: string;
    price: number;
    imageUrl: string;
}

export type item =  catalogItem | null