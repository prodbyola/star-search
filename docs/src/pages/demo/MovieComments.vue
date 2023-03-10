<template>
  <q-page class="row items-center justify-evenly">
    <h3 class="page-title">Comments on {{ movie?.name }}</h3>
    <q-card class="doc-card">
      <div class="empty-box">
        <q-icon class="eb-icon" name="comment" />
        <p v-if="!comments.length" class="eb-desc">Be the first to comment.</p>
      </div>
      <div class="comment-form">
        <q-input
          v-model="comment"
          class="comment-field"
          label="Write your comment here"
          outlined
        />
        <q-btn
          label="Submit"
          class="cf-action"
          @click="makeComment"
          :loading="submittingComment"
          flat
        />
      </div>
      <div v-if="comments.length" class="comment-list">
        <div
          v-for="(comment, index) in comments"
          :key="'comment-' + index"
          class="comment"
        >
          <div class="cm-head">
            <h3 class="ca">{{ comment.author }}</h3>
            <h3 class="ct">
              {{ new Date(comment.created_at).toLocaleString() }}
            </h3>
          </div>
          <p class="cmc">{{ comment.content }}</p>
        </div>
      </div>
    </q-card>
  </q-page>
</template>
<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue';
import { useStore, Comment } from 'stores/app';
import { api } from 'src/boot/axios';
import { ApiResponse } from 'src/shared';
import { useRoute } from 'vue-router';

const appStore = useStore();
const route = useRoute();
const movie = computed(() => appStore.currentMovie);
const comment = ref('');
const comments = ref<Comment[]>([]);
const movieID = parseInt(route.params.movieID as string);
const submittingComment = ref(false);

const makeComment = () => {
  submittingComment.value = true;
  const data = {
    movie_id: movieID,
    content: comment.value,
  };

  api.post('/add-comment', data).then((r) => {
    const resp = r.data as ApiResponse;
    if (!resp.error) comments.value.push(resp.message as Comment);
    submittingComment.value = false;
  });
};

onMounted(() => {
  api.get('/movie-comments/' + movieID).then((r) => {
    const resp = r.data as ApiResponse;
    if (!resp.error) comments.value = resp.message as Comment[];
    // console.log(comments.value);
  });
});
</script>
<style lang="scss" scoped>
@import './resp.scss';
.doc-card {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.comment-form {
  width: 60%;
  margin-bottom: 36px;

  .comment-field {
    width: 100%;
    margin-bottom: 8px;
  }
  .cf-action {
    background-color: $primary;
    color: white;
    float: right;
    width: 150px;
  }
}
.empty-box {
  @extend .doc-card;
  display: inline-flex;
  width: 100%;
  .eb-icon {
    font-size: 148px;
  }

  .eb-desc {
    font-size: 2.6rem;
  }
}

.comment-list {
  background-color: $secondary;
  color: white;
  width: 100%;
  padding: 12px 24px;
  
  .comment {
    margin-bottom: 24px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.61);
    padding: 12px 0;
      .cm-head {
        display: inline-flex;
        width: 36%;
        justify-content: space-between;
    
        h3 {
          font-size: 1rem;
          margin: 8px 0;
        }
      }
      .cmc {
        margin-left: 34px;
        font-size: 1.4rem;
      }
  }

}
</style>
<style lang="scss">
.comment-field.q-field--outlined .q-field__inner,
.comment-field.q-field--outlined .q-field__control:before,
.comment-field.q-field--outlined .q-field__control:after {
  height: 226px;
}
</style>
