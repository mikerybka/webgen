import { useState } from "react";

interface Field {
  meta: Meta;
  type: string;
}

interface FuncType {
  meta: Meta;
  inputs: Field[];
  outputs: Field[];
}

interface Type {
  kind: "primitive" | "struct" | "array" | "map" | "func";
  primitiveType?: string;
  structType?: {
    fields: Field[];
  };
  arrayType?: string;
  mapType?: string;
  funcType?: FuncType;
  methods: FuncType[];
}

interface Content {
  meta: Meta;
  type: Type;
  data: any;
}

interface Meta {
  location: string[];
  name: string;
  tags: {[key: string]: boolean};
  doc: string;
  comments: Comment[];
}

interface Comment {
  timestamp: number;
  author: string;
  text: string;
  replies: Comment[];
}

const defaultContent: Content | null = null;

function useContent() {
  const [content, setContent] = useState<Content |  null>(defaultContent);
  
  return content;
}

function App() {
  const content = useContent();
  if (content === null) {
    return <div className="text-red-500">Loading!!!...</div>;
  }

  return (
    <Navigation {...content.meta} >
      <Breadcrumbs {...content.meta} />
      <Name {...content.meta} />
      <Tags {...content.meta} />
      <Doc  {...content.meta} />
      <View {...content} />
      <Discussion {...content.meta} />
    </Navigation>
  );
}

export default App;

function Navigation(props: {location: string[]; children: any}) {
  return <div>{props.children}</div>
}

function Breadcrumbs(props: {location: string[]}) {
  return <div></div>
}

function Name(props: Meta) {
  return <div>Name</div>
}

function Tags(props: Meta) {
  return <div></div>
}

function Doc(props: Meta) {
  return <div></div>
}

function View(props: Content) {
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

function Discussion(props: Meta) {
  return <div></div>
}