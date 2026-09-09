interface catalogItem {
    Id: number;
    Name: string;
    Price: number;
    ImageUrl: string;
}

export type item =  catalogItem | null