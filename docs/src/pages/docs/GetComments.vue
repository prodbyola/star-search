<template>
  <ApiTemplate
    method="GET"
    name="4. Get Comments"
    desc="Gets comments associated with a movie. It accepts a single parameter 'movie_id'"
    :link="apiBase + '/movie-comments/:movie_id'"
  >
    <template v-slot:response>
      <div class="sig">
        <p class="resp-sig">
          message: <code class="light">[]Comment | string</code>
        </p>
        <p>
          Contains a list of <code>Comment</code> associated with if the movie_id if the request was
          successful. Please see the <code>message</code> section of <strong>3. Add Comment</strong> to see <code>Comment</code>' signature.
        </p>
      </div>
    </template>
  </ApiTemplate>
</template>
<script lang="ts" setup>
import { apiBase } from 'src/shared';
import ApiTemplate from './ApiTemplate.vue';
</script>
