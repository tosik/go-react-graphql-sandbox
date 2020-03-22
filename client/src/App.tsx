import React, { useState, useEffect } from 'react';
import './App.css';

import ApolloClient from "apollo-boost";
import gql from "graphql-tag";
import { ApolloProvider } from '@apollo/react-hooks';
import { useQuery, useMutation } from '@apollo/react-hooks';

const BOOKS_QUERY = gql`
query {
  books {
    title
    id
    price
    foo
  }
}
`

const CREATE_BOOK_MUTATION = gql`
mutation createBook {
  createBook(input : {
    title: "竹取物語",
    price: 1980,
    foo: {
      koreha: "nandemo",
      dekimasu: 9090909
    }
  }) {
    id
    title
    price
    foo
  }
}
`

const Books: React.FC = () => {
  var { loading, error, data } = useQuery(BOOKS_QUERY)

  if (loading) return (<p>Loading...</p>)
  if (error) return (<p>Error </p>)

  var list = []
  for (var item of data.books) {
    list.push(<li>{item.title}: {item.price}円</li>)
  }

  return (
    <div>
      <ul>{list}</ul>
    </div>
  )
}

const CreateBookButton: React.FC = () => {
  var [createBook, { data }] = useMutation(CREATE_BOOK_MUTATION)

  return (
    <button onClick={
      e => {
        e.preventDefault()
        createBook()
      }
    }> Create a book</button>
  )
}

const App: React.FC = () => {
  const client = new ApolloClient({
    uri: 'http://localhost:8080/query',
    request: operation => {
      operation.setContext({});
    }
  });

  return (
    <ApolloProvider client={client}>
      <div>
        <h2>My first Apollo app </h2>
        <Books />
        <CreateBookButton />
      </div>
    </ApolloProvider>
  );
}

export default App
