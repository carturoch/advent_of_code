defmodule InverseCaptcha do
  
  def sum(str) when is_binary(str) do
    [h|tail] = String.graphemes(str)
    sum([h|tail], h)
  end

  def sum([], _), do: 0
  def sum([a], a), do: String.to_integer(a)
  def sum([_a], _), do: 0
  def sum([a|[a|tail]], first), do: String.to_integer(a) + sum([a|tail], first)
  def sum([_a|[b|tail]], first), do: sum([b|tail], first)
end