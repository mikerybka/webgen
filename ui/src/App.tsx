function App() {
  return (
    <Navigation>
      <Breadcrumbs />
      <Name />
      <Tags />
      <Doc />
      <View />
      <Discussion />
    </Navigation>
  );
}

export default App;

function Navigation(props: {children: any}) {
  return <div>{props.children}</div>
}

function Breadcrumbs() {
  return <div></div>
}

function Name() {
  return <div>Name</div>
}

function Tags() {
  return <div></div>
}

function Doc() {
  return <div></div>
}

function View() {
  const generic = true;
  if (generic) {
    return <div>
      <JSONView />
      <Methods />
    </div>
  }
}

function JSONView() {
  return <div></div>
}

function Methods() {
  return <div></div>
}

function Discussion() {
  return <div></div>
}