<template>
  <v-form
    ref="form"
    lazy-validation
    v-model="formValid"
    v-if="isLoaded()"
  >
    <v-alert
      :value="formError"
      color="error"
      class="pb-2"
    >{{ formError }}</v-alert>

    <v-text-field
      v-model="item.id"
      :label="$t('id')"
      required
      :disabled="true"
      class="mb-4"
    ></v-text-field>

    <v-text-field
      v-model="item.name"
      :label="$t('name')"
      clearable
      required
      :disabled="formSaving"
      class="mb-4"
    ></v-text-field>

    <v-select
      v-model="item.project_id"
      :label="$t('projectName')"
      clearable
      :items="projects"
      item-value="id"
      item-text="name"
      :disabled="formSaving"
    ></v-select>

    <v-select
      clearable
      v-model="item.inventory_id"
      :label="$t('inventoryName')"
      :items="inventories"
      item-value="id"
      item-text="name"
      :disabled="formSaving"
      class="mb-4"
    ></v-select>

    <v-text-field
      v-model="item.max_parallel_tasks"
      :label="$t('maxParallelTasks')"
      required
      :disabled="formSaving"
      type="number"
      class="mb-4"
      oninput="if(this.value < 1) this.value = 1;"
    ></v-text-field>

  </v-form>
</template>

<script>
/* eslint-disable import/no-extraneous-dependencies,import/extensions */

import axios from 'axios';
import ItemFormBase from '@/components/ItemFormBase';

export default {
  mixins: [ItemFormBase],
  components: {
  },
  data() {
    return {
      projects: null,
      inventories: null,
    };
  },
  async created() {
    // Loading projects and inventories
    [this.projects, this.inventories] = (await Promise.all([
      await axios({
        keys: 'get',
        url: '/api/projects',
        responseType: 'json',
      }),
      await axios({
        keys: 'get',
        url: `/api/project/${this.projectId}/inventory`,
        responseType: 'json',
      }),
    ])).map((x) => x.data);
  },
  methods: {

    getSingleItemUrl() {
      return `/api/project/${this.projectId}/runner/${this.itemId}`;
    },

    isLoaded() {
      return this.item != null
        && this.projects != null
        && this.inventories != null;
    },

    beforeSave() {
      this.item.max_parallel_tasks = Number(this.item.max_parallel_tasks);
    },
  },
};
</script>
