import React from "react";
import { maxPts } from "../../consts";
import Checkbox from "./Checkbox";
import "./Scoreboard.css";

type Props = {
  resPts: number;
  spyPts: number;
};

const points = Array.from(Array(maxPts), (_, i) => i + 1);

export default ({ resPts, spyPts }: Props) => (
  <table className="Scoreboard">
    <tbody>
      <tr>
        <td>
          {points.map(k => (
            <Checkbox key={k} disabled checked={resPts >= k} />
          ))}
        </td>
        <td>
          <b>Resistance</b>
        </td>
      </tr>
    </tbody>
    <tr>
      <td>
        {points.map(k => (
          <Checkbox key={k} disabled checked={spyPts >= k} />
        ))}
      </td>
      <td>
        <b>Spies</b>
      </td>
    </tr>
  </table>
);
