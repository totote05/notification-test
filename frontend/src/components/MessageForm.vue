<template>
  <div>
    <h3 class="is-size-4">Send a notification</h3>
    <hr />
    <form @submit.prevent="onSubmit">
      <div class="field">
        <label class="label">Category</label>
        <div class="control">
          <category-selector v-model="category" />
        </div>
      </div>
      <div class="field">
        <label class="label">Message</label>
        <div class="control">
          <textarea
            class="textarea"
            :class="{ 'is-danger': messageRequired }"
            v-model="message"
          ></textarea>
        </div>
        <p v-if="messageRequired" class="help is-danger">
          This field is required
        </p>
      </div>
      <div class="field">
        <div class="control">
          <button
            class="button is-primary"
            :class="{ 'is-loading': sending }"
            type="submit"
            :disabled="messageRequired || sending"
          >
            Send
          </button>
        </div>
      </div>
    </form>
    <div
      v-if="!!sendErrorMessage"
      class="my-3 notification is-danger is-light has-text-centered"
    >
      We can't send the notification, reason:
      {{ sendErrorMessage }}
    </div>
  </div>
</template>

<script>
import { mapActions, mapGetters, mapMutations } from "vuex";
import CategorySelector from "./CategorySelector.vue";
export default {
  components: { CategorySelector },
  data() {
    return {
      category: "sport",
      message: "",
    };
  },
  computed: {
    ...mapGetters("message", ["sending", "sendErrorMessage"]),
    messageRequired() {
      return this.message === "";
    },
  },
  methods: {
    ...mapActions("message", ["sendMessage"]),
    ...mapMutations("message", ["setErrorMessage"]),
    onSubmit() {
      this.setErrorMessage("");

      if (!this.messageRequired) {
        this.sendMessage({ category: this.category, message: this.message });
        this.message = "";
        this.category = "sport";
      }
    },
  },
};
</script>
