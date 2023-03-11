<template>
    <q-page class="row items-center justify-evenly">
        <h3 class="page-title">{{ movie?.name }}' Characters</h3>
        
        <q-dialog v-model="sorters.open">
            <q-card class="sort-card">
                <h3>Sort & Filter</h3>
                <div class="sorters">
                    <q-select v-model="sorters.sort" :options="sortKeys" label="Sort By" outlined class="sf" />
                    <q-select v-model="sorters.order" :options="sortOrderKeys" label="Order By" class="sf" outlined />
                    <q-select v-model="sorters.filter" :options="filterKeys" label="Filter By" class="sf" outlined />
                    <q-btn label="Sort & Filter" icon="search" class="sa" @click="search" flat v-close-popup />
                </div>
            </q-card>
        </q-dialog>
        <q-spinner class="spinner" v-if="mounting" />

        <q-card v-if="results" class="doc-card">
            <div class="meta">
                <q-btn label="Sort & Filter" icon="search" class="search-btn" @click="sorters.open = !sorters.open" flat />
                <p>Total Count: {{ results.total_character_count }}</p>
                <p>Total Height: {{ results.total_character_height.ft }}</p>
            </div>
            <q-table
                :columns="columns"
                :rows="results?.characters"
                :pagination="{
                    rowsPerPage: 100
                }"
                :loading="loading"
            />
        </q-card>
    </q-page>
</template>
<script lang="ts" setup>
import { useStore } from 'src/stores/app';
import { computed, onMounted, reactive, ref } from 'vue';
import { Character } from 'src/stores/app';
import { useRoute } from 'vue-router';
import { api } from 'src/boot/axios';
import { ApiResponse } from 'src/shared';
import { QTableColumn } from 'quasar';

type QueryResult = {
    total_character_count: number
    total_character_height: {
        ft: string
        cm: number
    }
    characters: Character[]
}

const sortKeys = ['name', 'gender', 'height']
const sortOrderKeys = ['ascn', 'desc']
const filterKeys = ['male', 'female', 'hermaphrodite', 'n/a']

const sorters = reactive({
    open: false,
    sort: '',
    order: '',
    filter: '',
})

const loading = ref(false)
const mounting = ref(true)
const columns: QTableColumn[] = [
    {
        name: 'name',
        required: true,
        label: 'Name',
        align: 'left',
        field: (row: Character) => row.name,
        sortable: true
    },
    {
        name: 'gender',
        required: true,
        label: 'Gender',
        align: 'left',
        field: (row: Character) => row.gender,
        sortable: true
    },
    {
        name: 'height',
        required: true,
        label: 'Height',
        align: 'left',
        field: (row: Character) => row.height,
        sortable: true
    },
]

const appStore = useStore();
const movie = computed(() => appStore.currentMovie);
const route = useRoute()

const results = ref<QueryResult | undefined>()
const movieID = parseInt(route.params.movieID as string);
const baseUrl = 'movie-characters/'+movieID
const search = () => {
    const query = `${baseUrl}?sort_by=${sorters.sort}&sort_order=${sorters.order}&filter_by=${sorters.filter}`
    getResults(query)
}

const getResults = async(url: string) => {
    loading.value = true
    await api.get(url)
    .then(r => {
        const resp = r.data as ApiResponse
        if(!resp.error) results.value = resp.message as QueryResult
        loading.value = false
    })
}
onMounted(async() => {
    await getResults(baseUrl)
    mounting.value = false
})
</script>
<style lang="scss" scoped>
.sort-card {
    width: 500px;
    padding: 34px;

    h3 {
        margin: 4px 0;
        font-size: 2rem;
        text-align: center;
    }

    .sorters {
        width: 100%;

        .sf {
            margin-bottom: 8px;
        }

        .sa {
            @extend .search-btn;
            float: right;
        }
    }
}

.search-btn {
    margin-top: 8px;
    background-color: $primary;
    color: white;
}
</style>