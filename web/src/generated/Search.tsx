// Generated from screens/Search.aui; edit the AUI source.
import { Alert, Badge, Button, Card, Checkbox, Divider, Field, Heading, Icon, Input, Link, Page, Row, Section, Select, Stack, Text } from "@/components/ui"

export function Search({ data, onAction }: { data: any; onAction: (name: string) => void }) {
  return (
    <Page>
      <Heading level={1}>Find it in your Workspace</Heading>
      <Text>Search uses your own account in each app. Results are never collected with an administrator account.</Text>
      <Row>
        <Input name="query" onClick={() => onAction("search")} type="search" placeholder="Search files, notes, events, people and tasks…" />
        <Button onClick={() => onAction("search")}>Search</Button>
        <Select name="sourceFilter" options={data.sourceOptions} />
        <Select name="kindFilter" options={data.kindOptions} />
      </Row>
      <Row>
        <Field label="Modified from"><Input name="after" type="date" /></Field>
        <Field label="Modified through"><Input name="before" type="date" /></Field>
        <Field label="Tag (Tasks)"><Input name="tagFilter" /></Field>
        <Field label="Owner (when supplied by source)"><Input name="ownerFilter" /></Field>
      </Row>
      <Section id="search-summary" />
      {data.suggestions.map((item, i) => (
        <Button key={i} variant="secondary" item={item.term} onClick={() => onAction("suggestion")}>{item.term}</Button>
      ))}
      <Row>
        <Checkbox name="recent" label="Remember my recent searches (off by default)" />
        <Button variant="secondary" onClick={() => onAction("preferences")}>Save preference</Button>
        <Button variant="secondary" onClick={() => onAction("clear")}>Clear recent searches</Button>
      </Row>
      {data.recent.map((item, i) => (
        <Button key={i} variant="secondary" item={item.term} onClick={() => onAction("recent")}>{item.term}</Button>
      ))}
      {data.standalone && (
        <Alert>Open My Search through One Workspace to discover apps on the same origin. Standalone mode has no administrator crawler or default remote sources.</Alert>
      )}
      {data.statuses.map((item, i) => (
        <Row key={i}>
          <Text>{item.name}</Text>
          <Badge tone={item.tone}>{item.status}</Badge>
          {item.connect && (
            <Link href={item.path} external={true}>Open app to connect</Link>
          )}
        </Row>
      ))}
      {data.empty && (
        <Card variant="empty">
          <Icon name="Layers" />
          <Heading level={2}>{data.emptyTitle}</Heading>
          <Text>My Vault is excluded. Mail bodies and Chat messages are not indexed or searched. Each source enforces its own permissions.</Text>
        </Card>
      )}
      {data.groups.map((item, i) => (
        <Card key={i}>
          <Row>
            <Heading level={2}>{item.name}</Heading>
            <Badge>{item.count}</Badge>
          </Row>
          {item.results.map((item, i) => (
            <Stack key={i} gap="xs">
              <Link href={item.href}>{item.title}</Link>
              <Text variant="snippet">{item.snippet}</Text>
              <Text variant="caption">{item.detail}</Text>
              <Divider />
            </Stack>
          ))}
          {item.more && (
            <Button variant="secondary" item={item.id} onClick={() => onAction("more")}>Load more from this source</Button>
          )}
        </Card>
      ))}
      <Card>
        <Heading level={2}>Other installed apps</Heading>
        <Text variant="caption">These apps expose launch metadata only. Their private content is not searched. My Vault is always excluded.</Text>
        {data.unsupported.map((item, i) => (
          <Link key={i} href={item.href}>{item.name}</Link>
        ))}
      </Card>
    </Page>
  )
}
