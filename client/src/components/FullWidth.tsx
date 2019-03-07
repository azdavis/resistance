import React from "react";
import "./FullWidth.css";

type Props = {
  children: React.ReactNode;
};

export default ({ children }: Props): JSX.Element => (
  <div className="FullWidth">{children}</div>
);
