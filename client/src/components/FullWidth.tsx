import React from "react";
import "./FullWidth.css";

type Props = {
  children: React.ReactNode;
};

export default ({ children }: Props) => (
  <div className="FullWidth">{children}</div>
);
