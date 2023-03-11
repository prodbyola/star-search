<template>
  <q-page class="row items-center justify-evenly">
        <h3 v-if="!loading" class="page-title">Movie List</h3>
        <q-spinner class="spinner" v-if="loading" />
        <div class="movie-list" v-else>
            <q-card v-for="(movie, index) in movieList" :key="index" class="movie-box">
                    <h3 class="mb-name">{{ movie.name }}</h3>
                <p class="mb-op" v-html="movie.opening_crawl" />
                <div class="mb-actions">
                    <q-btn @click="getUrl(movie, 'movie-comments')" class="mba" :label="movie.comment_count + ' Comments'" flat />
                    <q-btn @click="getUrl(movie, 'movie-characters')" class="mba" label="Show Characters" flat />
                </div>
            </q-card>
        </div>
    </q-page>
</template>
<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { api } from 'src/boot/axios';
import { ApiResponse } from 'src/shared';
import { Movie, useStore } from 'stores/app'
import { useRouter } from 'vue-router';

const appStore = useStore()
const router = useRouter()
const loading = ref(true)

const movieList = ref<Movie[]>([])
const getUrl = (movie: Movie, name: string) => {
    appStore.currentMovie = movie
    router.push({
        name,
        params: {
            movieID: movie.id
        }
    })
}

onMounted(() => {
    api.get('/get-movies').then(r => {
        const resp: ApiResponse = r.data
        if(!resp.error){
            movieList.value = resp.message as Movie[]
            loading.value = false
        }
    })
})
</script>
<style lang="scss" scoped>
.movie-list {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    row-gap: 12px;
    width: 90%;

    .movie-box{
        width: 32%;
        padding: 16px;
        display: grid;

        .mb-name {
            margin: 20px 0;
            font-size: 2.4rem;
        }

        .mb-op {
            color: #5e5c5e;
        }

        .mb-actions {
            display: inline-flex;
            width: 100%;
            justify-content: space-between;
            align-items: center;
            color: #adadad;
            align-self: self-end;
        }
    }
}

body.screen--sm {
    .movie-box {
        width: 48%;
    }
}
body.screen--xs {
    .movie-list {
        width: 96%;
    }
    .movie-box {
        width: 100%;
    }
}
</style>