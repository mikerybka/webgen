import { useEffect, useState } from "react";

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

function useContent(location: string[] | null) {
  const [content, setContent] = useState<Content |  null>(defaultContent);
  useEffect(() => {
    if (location !== null) {
      fetchContent(location).then(setContent)
    }
  }, [JSON.stringify(location)])
  return content;
}

async function fetchContent(location: string[]) {
  const url = "https://api.brass.software/" + location.join("/")
  const res = await fetch(url)
  return await res.json();
}

function useLocation(): string[] | null {
  const [location, setLocation] = useState<string[] | null>(null);
  useEffect(() => {
    if (typeof window !== "undefined") {
      const loc = window.location.hostname + window.location.pathname;
      setLocation(loc.split("/").filter(Boolean))
    }
  }, []);
  return location;
}

function App() {
  const location = useLocation();
  const content = useContent(location);
  if (content === null) {
    return <div className="">Loading {JSON.stringify(location)}...</div>;
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