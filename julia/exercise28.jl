using Statistics

rationalize(successes, trials) = successes .// trials

probabilities(successes, trials) = successes ./ trials

function checkmean(successes, trials)
    r, p = mean(rationalize(successes, trials)), mean(probabilities(successes, trials))
    float(r) == p || r 
end

function checkprob(successes, trials)
    r, p = prod(rationalize(successes, trials)), prod(probabilities(successes, trials))
    float(r) == p || r
end