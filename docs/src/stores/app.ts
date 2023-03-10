import { defineStore } from 'pinia'

export type Movie = {
    id: number
    opening_crawl: string
    name: string
    comment_count: number
}

export type Comment = {
    id: number
    movie_id: number
    author: string
    content: string
    created_at: string
}

export type Character = {
    name: string,
    gender: string
    height: string
}

type AppStore = {
    currentMovie?: Movie
}

export const useStore = defineStore('appStore', {
    state: (): AppStore => ({
        currentMovie: undefined
    })
})