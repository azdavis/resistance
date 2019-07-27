import React from "react";
import "./ButtonLink.css";
import "./Button.css";
import "./Truncated.css";

type Props = {
  value: string;
  href: string;
};

export default ({ value, href }: Props) => (
  <a href={href} className="ButtonLink Button Button--enabled Truncated">
    {value}
  </a>
);
