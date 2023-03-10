<template>
    <ApiTemplate
      name="3. Add Comment"
      desc="Add comment to a movie."
      :link="apiBase + '/add-comment'"
      method="POST"
    >
      <template v-slot:body>
        <div class="sig">
          <p class="resp-sig">movie_id: <code class="light">number</code></p>
          <p>
            ID of the movie to which you want to add the comment. If the movie is
            not found or does not exist, we get the error message with error code
            <code>404</code>.
          </p>
        </div>
        <div class="sig">
          <p class="resp-sig">content: <code class="light">string</code></p>
          <p>
            Content of the comment you wish to add. Content <strong>must not exceed 500 characters</strong> or this could lead to a <code>417</code> error.
          </p>
        </div>
      </template>
      <template v-slot:response>
        <div class="sig">
          <p class="resp-sig">
            message: <code class="light">object | string</code>
          </p>
          <p>
            If request was successful, <code>message</code> contains information about the new comment (see below). Otherwise, it contains an error message.
          </p>
          <div class="ad-resp-indent">
            <div class="sig">
              <p class="resp-sig">
                id: <code class="light">number</code>
              </p>
              <p>ID of the newly created comment.</p>
            </div>
            <div class="sig">
              <p class="resp-sig">
                movie_id: <code class="light">number</code>
              </p>
              <p>ID of the movie to which the comment is associated.</p>
            </div>
            <div class="sig">
              <p class="resp-sig">
                author: <code class="light">string</code>
              </p>
              <p>Public IP address of the comment's author. This my might be in ipv4 or ipv6 format.</p>
            </div>
            <div class="sig">
              <p class="resp-sig">
                content: <code class="light">string</code>
              </p>
              <p>Content of the comment.</p>
            </div>
            <div class="sig">
              <p class="resp-sig">
                created_at: <code class="light">string</code>
              </p>
              <p>Date the comment was created in UTC format.</p>
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
  