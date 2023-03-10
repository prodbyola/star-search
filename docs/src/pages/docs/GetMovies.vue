<template>
    <ApiTemplate method="GET" name="1. Get Movies" desc="Lists all available movies. It accepts no url or query parameters." :link="apiBase+'/get-movies'">
        <template v-slot:response>
            <div class="sig">
                <p class="resp-sig">message: <code class="light">[]Movie | string</code></p>
                <p>
                    Contains a list of <code>Movie</code> (see below) if request was successful or contains error message if the request fails.
                </p>
                <div class="ad-resp-indent">
                    <div class="sig">
                        <p class="resp-sig">Movie: <code class="light">object</code></p>
                        <p>
                            An object containing information about the movie. It contains the following properties.
                        </p>
                    </div>
                    <div class="ad-resp-indent">
                    <div class="sig">
                        <p class="resp-sig">id: <code class="light">number</code></p>
                        <p>
                            ID of the movie.
                        </p>
                    </div>
                    <div class="sig">
                        <p class="resp-sig">name: <code class="light">string</code></p>
                        <p>
                            Name of the movie.
                        </p>
                    </div>
                    <div class="sig">
                        <p class="resp-sig">opening_crawl: <code class="light">string</code></p>
                        <p>
                            The movie's opening crawl.
                        </p>
                    </div>
                    <div class="sig">
                        <p class="resp-sig">comment_count: <code class="light">number</code></p>
                        <p>
                            Total number of comments on the movie.
                        </p>
                    </div>
                </div>
                </div>
            </div>
        </template>
    </ApiTemplate>
</template>
<script lang="ts" setup>
import { apiBase } from 'src/shared';
import ApiTemplate from './ApiTemplate.vue';
</script>