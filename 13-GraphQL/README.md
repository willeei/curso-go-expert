# Queries
```
mutation createCategory {
  createCategory(input: {
    name: "Tecnologia",
    description: "Curso"
  }) {
    id
    name
    description
  }
}

mutation creatCourse {
  createCourse(input: {
    name: "FullCycle"
    description: "Muito bom"
    categoryId: "d1243b3e-ad67-4fba-9a58-26a6247d05a9"
  }) {
    id
    name
    description
  }
}

query queryCategories {
  categories {
    id
    name
    description
  }
}

query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
    }
  }
}

query queryCourses {
  courses {
    id
    name
    description
  }
}

query queryCourseWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}

```

# Doc
https://gqlgen.com/