<template>
  <ApiTemplate
    name="2. Get Characters"
    desc="Lists all characters casted in a movie. It accepts one url parameter 'movied_id', which is the id of movie. It accepts additonal query parameters (see queries below)."
    :link="apiBase + '/:movie_id?sort_by=name&sort_order=ascn&filter_by=female'"
    has-params
    has-queries
    method="GET"
  >
    <template v-slot:parameters>
      <div class="sig">
        <p class="resp-sig">movie_id: <code class="light">number</code></p>
        <p>
          ID of the movie whose characters are to be retrieved. If the movie is
          not found or does not exist, we get the error message with error code
          <code>404</code>.
        </p>
      </div>
    </template>
    <template v-slot:queries>
      <div class="sig">
        <p class="resp-sig">
          sort_by: <code class="light">name | gender | height</code>
        </p>
        <p>
          Optional query parameter that tells how you want the characters to be
          sorted. This can either be <code>name</code>, <code>gender</code> or
          <code>height</code>. If not provided, the characters will be retrieved
          as received from swapi.
        </p>
      </div>
      <div class="sig">
        <p class="resp-sig">
          sort_order: <code class="light">ascn | desc</code>
        </p>
        <p>
          Optional query parameter that tells what order the characters'
          retrieval should follow according to the
          <code>sort_by</code> parameter. This can either be
          <code>ascn</code> to mean ascending or <code>desc</code> to mean
          descending order. If not provided, the characters will be retrieved as
          received from swapi.
        </p>
      </div>
      <div class="sig">
        <p class="resp-sig">
          filter_by:
          <code class="light">male | female | hermaphrodite | n/a</code>
        </p>
        <p>
          Optional query parameter that filters characters according to their
          gender. If not provided, the characters will be retrieved as received
          from swapi.
        </p>
      </div>
    </template>
    <template v-slot:response>
      <div class="sig">
        <p class="resp-sig">
          message: <code class="light">object | string</code>
        </p>
        <p>
          If request was successful, <code>message</code> contains information about retrieved
          characters (see below). Otherwise, it contains an error message.
        </p>
        <div class="ad-resp-indent">
          <div class="sig">
            <p class="resp-sig">
              total_character_count: <code class="light">number</code>
            </p>
            <p>Total number of characters received from the request.</p>
          </div>
          <div class="sig">
            <p class="resp-sig">
              total_character_height: <code class="light">object</code>
            </p>
            <p>
              Total calculation of height of all characters received. This is an
              object with members <code>ft</code> and <code>cm</code>. Use
              <code>ft</code> to retrieve the height as feet and inches, e.g
              <code>100ft 7.09in</code>. Use <code>cm</code> to retrieve the
              height in centimeters (number).
            </p>
          </div>
          <div class="sig">
            <p class="resp-sig">
              characters: <code class="light">[]object</code>
            </p>
            <p>A list of all the characters received.</p>
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
