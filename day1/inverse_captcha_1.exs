defmodule InverseCaptcha do
  
  def sum1(str), do: calculate(str, 1)
  def sum2(str), do: calculate(str, div(String.length(str), 2))
  
  def calculate(str, offset) do
    map = 0..(String.length(str) - 1)
    |> Enum.zip(String.graphemes(str)) 
    |> Enum.into(%{})

    map
    |> Enum.filter(fn {k,_v} ->
      is_match?(map, k, offset)
    end)
    |> Enum.into(%{})
    |> Map.values
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum
  end

  def is_match?(map, idx, offset) do
    ref_idx = cond do
      idx + offset >= Enum.count(map) ->
        idx + offset - Enum.count(map)
      true ->
        idx + offset
    end
    map[idx] == map[ref_idx]
  end
end