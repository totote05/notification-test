<template>
  <div class="select">
    <select :value="value" @input="onChange">
      <option :key="value" v-for="{ value, label } in options" :value="value">
        {{ label }}
      </option>
    </select>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    value: String,
  },
  emit: ["input"],
  data() {
    return {
      selected: "sport",
      options: [
        { value: "sport", label: "Sports" },
        { value: "movies", label: "Movies" },
        { value: "finance", label: "Finance" },
      ],
    };
  },
  methods: {
    onChange(e) {
      this.$emit("input", e.target.value);
    },
  },
};
</script>
