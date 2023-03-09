<template>
    <ApiTemplate name="1. Get Movies" desc="Lists all available movies. It accepts no url or query parameters." :link="apiBase+'/get-movies'">
                <template v-slot:response>
                    <p class="adr-head">Response:</p>
                    <div class="ad-resp-indent">
                        <div class="sig">
                            <p class="resp-sig">error: <code class="light">boolean</code></p>
                            <p>
                                Indicates whether the request was successful or not. If <code>true</code>, it means the request is not successful because of an error. The error message
                                would be available in the <code>message</code> property of the response. If <code>false</code>, it means the request is successful and the resource would be 
                                <code>message</code> property of the response.
                            </p>
                        </div>
                        <div class="sig">
                            <p class="resp-sig">code: <code class="light">number</code></p>
                            <p>
                                The response code. <code>200</code> usually means the request is successful. If there's an error, the error code would be found here.
                            </p>
                        </div>
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
                    </div>
                </template>
            </ApiTemplate>
</template>
<script lang="ts" setup>
import { apiBase } from 'src/shared';
import ApiTemplate from './ApiTemplate.vue';
</script>